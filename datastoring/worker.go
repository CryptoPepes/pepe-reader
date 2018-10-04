package datastoring

import (
	"cloud.google.com/go/datastore"
	"cryptopepe.io/cryptopepe-reader/reader"
	"cryptopepe.io/cryptopepe-svg/builder"
	"github.com/rakanalh/scheduler"
	scheduleStorage "github.com/rakanalh/scheduler/storage"
	"cryptopepe.io/cryptopepe-reader/datastoring/events"
	"cryptopepe.io/cryptopepe-reader/datastoring/triggers"
	"cryptopepe.io/cryptopepe-reader/datastoring/data"
	"cryptopepe.io/cryptopepe-reader/datastoring/pub"
	"time"
	"log"
	"cryptopepe.io/cryptopepe-reader/datastoring/sub"
	"cryptopepe.io/cryptopepe-reader/datastoring/sub/event"
	"cryptopepe.io/cryptopepe-reader/bio-gen"
)

type BackfillTask struct {
	FromBlock uint64
	ToBlock   uint64
}

type Worker struct {
	reader reader.Reader
	dc     *datastore.Client

	svgBuilder *builder.SVGBuilder
	bioGenerator *bio_gen.BioGenerator

	eventHub   *events.EventHub
	triggerHub *triggers.TriggerHub

	entityBuf *data.EntityBuffer

	sched scheduler.Scheduler

	backfillTasks chan BackfillTask
}

func NewPepeDataWorker(r reader.Reader, dc *datastore.Client) *Worker {

	worker := new(Worker)
	worker.reader = r
	worker.dc = dc

	worker.svgBuilder = new(builder.SVGBuilder)
	worker.svgBuilder.Load()

	worker.bioGenerator = new (bio_gen.BioGenerator)
	worker.bioGenerator.Load()

	worker.eventHub = events.NewEventHub()
	worker.triggerHub = triggers.NewTriggerHub()
	worker.triggerHub.TriggerMap(worker.eventHub)

	worker.entityBuf = data.NewEntityBuffer(dc)

	worker.backfillTasks = make(chan BackfillTask, 10)

	return worker
}

func (worker *Worker) Close() {
	worker.sched.Stop()
}

func (worker *Worker) startBackfillExecutor() {

	for {
		task := <-worker.backfillTasks

		log.Printf("Running backfill for block #%d - #%d.\n", task.FromBlock, task.ToBlock)
		pub.StartBackfills(worker.eventHub, worker.reader, task.FromBlock, task.ToBlock)

		// Wait till the entity buffer has room for more backfills
		worker.entityBuf.WaitReady()
	}
}

func (worker *Worker) StartSchedule(runBackfills bool) {
	store := scheduleStorage.NewNoOpStorage()
	worker.sched = scheduler.New(store)

	evCtx := &event.EventContext{
		Reader:     worker.reader,
		EntityBuf:  worker.entityBuf,
		SvgBuilder: worker.svgBuilder,
		BioGenerator: worker.bioGenerator,
	}

	// Make sure that all triggers will be handled by the subscribers.
	sub.HandleAll(worker.triggerHub, evCtx)

	if runBackfills {
		worker.runBackfillSchedule()

		go worker.startBackfillExecutor()
	}

	watchSubs := pub.StartWatchers(worker.eventHub, worker.reader)

	defer func() {
		for i, watcher := range watchSubs {
			log.Printf("Closing watcher %d\n", i)
			watcher.Unsubscribe()
		}
	}()

	// Every minute: try to empty buffer, update the database
	worker.sched.RunEvery(time.Minute, func() {
		worker.entityBuf.UpdateAll()
	})

	// Update the pepes data every 5 min. (Remove old auction data etc.)
	worker.sched.RunEvery(time.Minute*5, func() {
		worker.entityBuf.UpdatePepesData()
	})

	worker.sched.Start()
	worker.sched.Wait()
}

const contractBaseBlock uint64 = 3327940

func (worker *Worker) runBackfillSchedule() {

	// Backfill util
	//-------------------------------------------------------------------------

	//backfill in windows of at most 15000 blocks (~60 hours)
	backfillWindowSize := uint64(15000)

	runBackFillWindow := func(blocks uint64) {
		var err error
		var currentBlock uint64
		currentBlock, err = worker.reader.GetCurrentBlock()
		if err != nil {
			log.Println("Failed recent-only backfill, could not retrieve current block number!")
			log.Println(err)
			return
		}

		// 0 is special case, full backfill.
		if blocks == 0 {
			// Contract created in block 3327965
			blocks = currentBlock - contractBaseBlock
		}

		startPoint := currentBlock - blocks

		windows := (blocks / backfillWindowSize) + 1

		log.Printf("Running backfill for last %d blocks, starting from %d, chunked up in %d windows.\n", blocks, startPoint, windows)

		for i := uint64(0); i < windows; i++ {
			from := startPoint + (i * backfillWindowSize)
			to := from + backfillWindowSize
			// clip last window
			if to > currentBlock {
				to = currentBlock
			}

			worker.backfillTasks <- BackfillTask{FromBlock: from, ToBlock: to}
		}

	}


	// Initial run
	//-------------------------------------------------------------------------

	// Start asynchronous full backfill
	go runBackFillWindow(uint64(0))

	// Schedule
	//-------------------------------------------------------------------------
	// Assuming a block time of ~14 sec.

	// Quick backfills for quick consistency fixes.


	// Partial Backfill every 5 mins (25*14/60 = 5.8 -> 5)
	worker.sched.RunEvery(5*time.Minute, runBackFillWindow, uint64(25))

	// Partial Backfill every 25 mins (120*14/60 = 28 -> 25)
	worker.sched.RunEvery(25*time.Minute, runBackFillWindow, uint64(120))

	// Partial Backfill every 55 mins (240*14/60= 56 -> 55),
	//  window overlap to cover for faster eventual consistency.
	worker.sched.RunEvery(55*time.Minute, runBackFillWindow, uint64(240))

	// Full Backfill every 360 mins
	worker.sched.RunEvery(360*time.Minute, runBackFillWindow, uint64(0))

}
