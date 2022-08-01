import kotlinx.coroutines.*

fun main() = runBlocking {
    val begin = System.nanoTime()
    suspend fun networkCall(i:Int) {
        delay(1000)
        println(i)
    }

    println("start")

    runBlocking {
        // launch always need a CoroutineScope to run in
        for (i in 1..1000000) launch { networkCall(i) }
        println("called task1 and task2 from ${Thread.currentThread()}")
    }

    println("done")
    val end = System.nanoTime()
    println("Elapsed time in ms: ${(end-begin)/ 1_000_000}")
}