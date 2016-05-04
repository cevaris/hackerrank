package com.cevaris.hackerrank.challenges

/**
 * https://www.hackerrank.com/challenges/solve-me-second
 */
object SolveMeSecond extends App {

  val n = readInt
  val lines = (1 to n).map((i) => readLine)

  def func(line: String) =
    line.split(" ").map(_.toInt).sum

  lines.map(func)
    .foreach(println)
}

object SolveMeSecondTest extends App {

  import SolveMeSecond._

  assert(func("33 434") == 467)
  assert(func("0 33") == 33)
  assert(func("-13 33") == 20)
}