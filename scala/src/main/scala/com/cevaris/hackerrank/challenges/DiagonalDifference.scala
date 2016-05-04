package com.cevaris.hackerrank.challenges

/**
 * https://www.hackerrank.com/challenges/diagonal-difference
 */
object DiagonalDifference extends App {
  val n = scala.io.StdIn.readInt()
  val lines = (1 to n).map((i) => scala.io.StdIn.readLine())
  val matrix: Seq[Seq[Int]] = lines.map(_.split(" ").map(_.toInt).toSeq).toSeq

  def diagonal(m: Seq[Seq[Int]]) = m.indices map ((i) => m(i)(i))

  def func(m: Seq[Seq[Int]]) = {
    val sliceA = diagonal(m)
    val sliceB = diagonal(m.reverse)
    Math.abs(sliceA.sum - sliceB.sum)
  }

  println(func(matrix))
}

object DiagonalDifferenceTest extends App {

  assert(DiagonalDifference.func(Seq(Seq(1, 2), Seq(1, 2))) == 0)
  assert(DiagonalDifference.func(Seq(Seq(11, 2, 4), Seq(4, 5, 6), Seq(10, 8, -12))) == 15)
}