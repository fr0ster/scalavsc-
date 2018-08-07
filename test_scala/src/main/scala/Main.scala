//import scala.collection.parallel._

object Main extends App {
  //println("Hello, World!")
  val toval = 2147483647 - 1
  var summa = BigInt(0)
  (1 to toval).foreach(x => summa += x)
  //println(s"Summa from 1 to ${toval} = ${sum}")
  //val xs = (BigInt(1) to BigInt(toval)).toParArray
  //xs.tasksupport = new ForkJoinTaskSupport(new scala.concurrent.forkjoin.ForkJoinPool(2))
  //val summa = xs.fold(BigInt(0)){_+_}
  println(s"Summa from 1 to ${toval} = ${ summa }")
}
