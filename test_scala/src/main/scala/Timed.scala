import scala.util.Random

object Timed {
    @volatile var dummy: Any = _
    def apply[T](body: =>T): (Double, Any) = {
        val start = System.nanoTime
        dummy = body
        val end = System.nanoTime
        (((end - start) / 1000) / 1000.0, dummy)
    }
}

