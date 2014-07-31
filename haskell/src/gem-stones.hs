#!/usr/bin/env runhaskell

{-

John has discovered various rocks. Each rock is composed of various elements, and each element is represented by a lowercase latin letter from 'a' to 'z'. An element can be present multiple times in a rock. An element is called a 'gem-element' if it occurs at least once in each of the rocks.

Given the list of rocks with their compositions, display the number of gem-elements that exist in those rocks.

Input Format
The first line consists of N, the number of rocks.
Each of the next N lines contain rocks' composition. Each composition consists of lowercase letters of English alphabet.

Output Format
Print the number gem-elements that exist in those rocks.

Constraints
1 ≤ N ≤ 100
Each composition consists of only small latin letters ('a'-'z').
1 ≤ Length of each composition ≤ 100

Sample Input

3
abcdde
baccd
eeabg

Sample Output

2

Explanation
Only "a", "b" are the two kind of gem-elements, since these are the only characters that occur in each of the rocks' composition.

-}

import Control.Monad
--import Data.Map
--import qualified Data.Map as Map
--import Data.List.Split 
import Debug.Trace
import Data.List


process s = let (source, samples) = sampleGemStone $ map dedup s
            in countGemStones source samples 0

dedup :: (Ord a, Eq a) => [a] -> [a]
dedup s = sort $ nub s

sampleGemStone :: [[Char]] -> ([Char], [[Char]])
sampleGemStone ls = (head ls, tail ls)

analyzeGemSample :: Char -> [[Char]] -> Int
analyzeGemSample c ls = length $ filter (\x -> elem c x) ls
--analyzeGemSample c ls = foldl (\s acc -> if elem c s then 1 else 0) ls 0
--analyzeGemSample c [] n     = n
--analyzeGemSample c (x:xs) n = if elem c x
--                              then analyzeGemSample c xs (n+1)
--                              else analyzeGemSample c xs (n)

-- | trace (show x ++ " " ++ show ls ++ " " ++ show acc) True 
countGemStones [] _ acc              = acc
countGemStones (x:xs) ls acc = if analyzeGemSample x ls == length ls
                               then countGemStones xs ls (acc+1)
                               else countGemStones xs ls acc




--freq s = map (\x -> (head x, length x)) $ group $ sort "happy"
--freqList :: [Char] -> [(Char, Integer)]
--freqList s = Map.toList $ Map.fromListWith (+) [(c, 1) | c <- s]


main = do
    testCount <- getLine
    samples <- replicateM (read testCount) getLine
    print $ process samples
    --mapM_ print (process samples)

    --mapM_ print ([ process s | s <- samples])
