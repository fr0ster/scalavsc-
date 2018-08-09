package jmh.main

// Must not be in default package
import java.util.concurrent.TimeUnit

import org.openjdk.jmh.annotations.{Benchmark, BenchmarkMode, Mode, OutputTimeUnit}

import scala.util.Random

object xsVector {
  val r = new scala.util.Random(100)
  val xs = for (i <- 1 to 5000000) yield r.nextInt(100)
  val xs_par = xs.par
}

/* Default settings for benchmarks in this class */
@OutputTimeUnit(TimeUnit.MILLISECONDS)
@BenchmarkMode(Array(Mode.Throughput))
class TestParVector {

  def toVectorFilter(xs: Seq[Int]) = xs.filter(_%3==0)
  def toParVectorFilter(xs: scala.collection.parallel.immutable.ParSeq[Int]) = xs.filter(_%3==0)
  @Benchmark
  def filtering: Unit = toVectorFilter(xsVector.xs)
  @Benchmark
  def parfiltering: Unit = toParVectorFilter(xsVector.xs_par)

  //@Benchmark
  //def interpolation: Unit = toHexStringInterp(randomArray)

  //@Benchmark
  //def format: Unit = toHexStringFormat(randomArray)


  //@Benchmark
  //def stringManip: Unit = toHexString(randomArray)

  //def toHexStringInterp(bytes: Array[Byte]) =
  //  bytes.map(b => f"$b%02x").mkString

  /*def toHexString(bytes: Array[Byte]) = {
    val hexArray: Array[Byte] = Array(
      '0', '1', '2', '3', '4',
      '5', '6', '7', '8', '9',
      'A', 'B', 'C', 'D', 'E',
      'F')
    val hexChars = Array.fill(bytes.size * 2)(0.toByte)
    for {
      j <- 0 to bytes.length - 1
      v = bytes(j) & 0xFF
    } {
      hexChars(j * 2) = hexArray(v >>> 4)
      hexChars(j * 2 + 1) = hexArray(v & 0x0F)
    }
    new String(hexChars)
  }

  def toHexStringFormat(bytes: Array[Byte]) =
    bytes.map(b => "%02x".format(b)).mkString


  def randomArray: Array[Byte] = {
    val a = Array.fill(20)(0.toByte)
    scala.util.Random.nextBytes(a)
    a
  }*/

}
