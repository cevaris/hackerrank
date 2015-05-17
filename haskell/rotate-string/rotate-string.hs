#!/usr/bin/env runhaskell
import Data.List

helper :: [String] -> Int -> String -> [String]
helper acc 0 _ = acc
helper acc step (x:xs) = helper (acc ++ [ls]) (step-1) ls
    where
      ls = xs ++ [x]


solve :: String -> String
solve s = unwords $ helper [] (length s) s

main :: IO ()
main = do
  n <- getLine
  input <- sequence (replicate (read n :: Int) getLine)
  mapM_ (putStrLn . solve) input
