import Control.Monad

-- Taken from
-- http://www.quora.com/What-is-the-most-efficient-algorithm-to-check-if-a-number-is-a-Fibonacci-Number

sqrt5 = sqrt 5

goldenRatio :: Double
goldenRatio = (1 + sqrt5) / 2

fibN :: Double -> Int
fibN n = floor $ (goldenRatio ** n) / sqrt5 + 0.5

fibStepN :: Double -> Int
fibStepN n = floor $ logBase goldenRatio (n * sqrt5) + 0.5

solve :: Int -> String
solve n = if n == (fibN . fromIntegral .  fibStepN . fromIntegral) n
          then "IsFibo"
          else "IsNotFibo"

main :: IO ()
main = do
  n <- getLine
  input <- replicateM (read n) getLine
  mapM_ (putStrLn . solve . read) input
