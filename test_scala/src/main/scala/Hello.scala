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
/*object Timed {
    @volatile var dummy: Any = _
    def apply[T](body: =>T): Double = {
        val start = System.nanoTime
        dummy = body
        val end = System.nanoTime
        ((end - start) / 1000) / 1000.0
    }
}*/

object Hello extends App {
  println("Hello, World!")
  val timed = Timed({(1L to Int.MaxValue).fold(0L)(_+_)})
  println(s"time execution long loop == ${timed._1}")
  println(s"result execution long loop == \n${timed._2}")
}
