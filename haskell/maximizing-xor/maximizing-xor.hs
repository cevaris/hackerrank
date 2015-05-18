import Data.Bits

solve :: Int -> Int -> Int
solve l r = maximum [ xor i j | i <- [l..r], j <- [i..r]]

main :: IO ()
main = do
  l <- readLn
  r <- readLn
  print $ solve l r
