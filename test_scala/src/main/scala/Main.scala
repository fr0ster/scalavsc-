object Main extends App {
  //println("Hello, World!")
  //val toval = 2147483647L - 1L
  val toval: Long = Int.MaxValue / 100
  //var summa = BigInt(0)
  var summa = 0L
  (1L to toval).foreach(x => summa += x)
  println(s"Summa from 1 to ${toval} = ${ summa }")
}
