object ParMain extends App {
  //println("Hello, World!")
  //val toval = 2147483647L - 1L
  val toval: Long = Int.MaxValue / 100
  //var summa = BigInt(0)
  val summa = (1L to toval).par.fold(0L)(_+_)
  println(s"Summa from 1 to ${toval} = ${ summa }")
}
