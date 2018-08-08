import scala.util.Random

object Hello extends App {
  println("Hello, World!")
  //val timed = Timed({(1L to Int.MaxValue).fold(0L)(_+_)})
  val timed = Timed({(1L to Int.MaxValue).par.fold(0L)(_+_)})
  println(s"time execution long loop == ${timed._1}")
  println(s"result execution long loop == \n${timed._2}")
}
