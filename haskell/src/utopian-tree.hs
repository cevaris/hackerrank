#!/usr/bin/env runhaskell
import Control.Monad

--The Utopian tree goes through 2 cycles of growth every year. The first growth cycle of the tree occurs during the monsoon, when it doubles in height. The second growth cycle of the tree occurs during the summer, when its height increases by 1 meter.
--Now, a new Utopian tree sapling is planted at the onset of the monsoon. Its height is 1 meter. Can you find the height of the tree after N growth cycles?

--Input Format
--The first line contains an integer, T, the number of test cases.
--T lines follow. Each line contains an integer, N, that denotes the number of cycles for that test case.

--Constraints
--1 <= T <= 10
--0 <= N <= 60

--Output Format
--For each test case, print the height of the Utopian tree after N cycles.

--Sample Input #00:

--2
--0
--1

--Sample Output #00:

--1
--2

--Explanation #00:
--There are 2 test cases. When N = 0, the height of the tree remains unchanged. When N = 1, the tree doubles its height as it's planted just before the onset of monsoon.

--Sample Input: #01:

--2
--3
--4

--Sample Output: #01:

--6
--7

--Explanation: #01:
--There are 2 testcases.
--N = 3:
--the height of the tree at the end of the 1st cycle = 2
--the height of the tree at the end of the 2nd cycle = 3
--the height of the tree at the end of the 3rd cycle = 6

--N = 4:
--the height of the tree at the end of the 4th cycle = 7
tree :: Char -> Integer -> Integer -> Integer
tree _ height 0 = height
tree 'M' height cycle = tree 'S' (height+1) (cycle-1)
tree 'S' height cycle = tree 'M' (height*2) (cycle-1)

--utopia :: Integer -> [Integer] -> [Integer] -> IO ()
--utopia 0 x = x
--utopia testCount x = do
--                    line <- getLine
--                    print (tree 'S' 1 (read line :: Integer))
--                    utopia (testCount-1) x

get2Lines = do
    line1 <- getLine
    line2 <- getLine
    return ((read line1 :: Integer),(read line2:: Integer))


--readInt :: () -> Integer
readInt = do
        line <- getLine
        return (read line :: Integer)

--readTests :: Integer -> [String] -> [String]
--readTests 0 xs = xs
--readTests testCount xs = do 
--                    readTests (testCount-1) ((lift getLine):xs)



--readLines :: IO [String]
--readLines = do  c <- getLine
--                if c == "" 
--                    then return []
--                    else do s <- readLines
--                            return (c:s)

--readInts :: IO [Integer]
--readInts = do   c <- getLine
--                if c == "" 
--                    then return []
--                    else do s <- readInts
--                            return ((read c :: Integer):s)

--readInts :: Integer -> Integer
--readInts 0 = []
--readInts x = do
--            c <- getLine
--            s <- readInts (x-1)
--            return ((read c :: Integer):s)


main = do
    testCount <- getLine
    inputs <- replicateM (read testCount) readInt
    print inputs
    --hieghts <- readInts
    --print hieghts
    --lines <- readLines
    --let hieghts = readTests (read testCount :: Integer) []
    --return hieghts
    --print 32
    --print x
    --let testCount' = (read testCount :: Integer)
    --    tests = readTests testCount'  []
    --print 10
    --utopia (read testCount :: Integer) []















