#!/usr/bin/env runhaskell
import Data.Char

data FunnyResult = Funny | NotFunny

instance Show FunnyResult where
    show = printFunny

printFunny :: FunnyResult -> String
printFunny Funny    = "Funny"
printFunny NotFunny = "Not Funny"

convert :: String -> [Int]
convert s = map ord s

helper :: (Num a, Eq a) => [Bool] -> [a] -> [a] -> [Bool]
helper acc [] []             = acc
helper acc (_:[]) (_:[])     = acc
helper acc (a:b:cs) (x:y:zs) = helper (acc ++ [constraint ((a,b), (x,y))]) cs zs

constraint :: (Num a, Eq a) => ((a, a), (a, a)) -> Bool
constraint ((x,y), (j,k)) = (abs (x - y) == abs (j - k))

solve :: String -> FunnyResult
solve s = result where
    a      = convert s
    b      = reverse a
    result = case False `elem` (helper [] a b) of
               True  -> NotFunny
               False -> Funny

main :: IO ()
main = do
  n <- getLine
  input <- sequence (replicate (read n :: Int) getLine)
  mapM_ (print . solve) input
