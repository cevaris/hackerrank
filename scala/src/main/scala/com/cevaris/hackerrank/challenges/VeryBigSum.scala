package com.cevaris.hackerrank.challenges

/**
 * https://www.hackerrank.com/challenges/a-very-big-sum
 */
object VeryBigSum extends App {

  val n = readInt
  val lines = Seq(readLine())

  def func(line: String) =
    line.split(" ").map(_.toLong).sum

  lines.map(func)
    .foreach(println)
}

object VeryBigSumTest extends App {

  import VeryBigSum._

  assert(func("2147483647 2147483647") == 4294967294L)
  assert(func("33 434") == 467)
  assert(func("0 33") == 33)
  assert(func("-13 33") == 20)
}