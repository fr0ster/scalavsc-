import scala.util.Random

object ParTest extends App {
  println("Hello, World!")
  val r = new scala.util.Random(100)
  val xs = for (i <- 1 to 500000) yield r.nextInt(100)
  val xs_par = xs.par
  println(s"time execution Map for Vector == ${Vector.range(0,1000).map(x=>Timed(xs.filter(_%3==0))._1).fold(0.0){_+_}/1000}")
  println(s"result execution Map for ParVector == ${Vector.range(0,1000).map(x=>Timed(xs_par.filter(_%3==0))._1).fold(0.0){_+_}/1000}")
}
