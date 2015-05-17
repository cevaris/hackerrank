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

pair :: [a] -> [(a, a)] -> [(a, a)]
pair [] acc       = acc
pair (_:[]) acc   = acc
pair (x:y:ls) acc = pair ([y] ++ ls) (acc ++ [(x, y)])

constraint :: ((Int, Int), (Int, Int)) -> Bool
constraint ((x,y), (j,k)) = (abs (x - y) == abs (j - k))

solve :: String -> FunnyResult
solve s = let x =  convert s
              y = zip (pair x []) (pair (reverse x) [])
              z = False `elem` (map constraint y)
              r = case z of
                    True  -> NotFunny
                    False -> Funny
          in r

main :: IO ()
main = do
  n <- getLine
  input <- sequence (replicate (read n :: Int) getLine)
  mapM_ (print . solve) input
