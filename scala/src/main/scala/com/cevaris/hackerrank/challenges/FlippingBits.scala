package com.cevaris.hackerrank.challenges

/**
 * https://www.hackerrank.com/challenges/flipping-bits
 */
object FlippingBits extends App {
  val n = readInt()
  val lines = (1 to n).map((i) => readLine())
  val data = lines.map(_.toLong)

  def getBinaryString(x: Long): String =
    String.format("%32s", x.toBinaryString).replace(' ', '0')

  def func(x: Long): Long = {
    java.lang.Long.parseLong(
      getBinaryString(x).map((c) => if (c == '0') '1' else '0'),
      2
    )
  }

  data.map(func).foreach(println)
}

object FlippingBitsTest extends App {
  assert(FlippingBits.func(2147483647) == 2147483648L)
  assert(FlippingBits.func(1) == 4294967294L)
  assert(FlippingBits.func(0) == 4294967295L)
}