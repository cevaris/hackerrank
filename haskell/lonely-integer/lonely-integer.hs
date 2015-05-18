import Control.Monad
import Data.List
import Data.Ord (comparing)

solve :: String -> String
solve = head . join . sortBy (comparing length) . group . sort . words

main :: IO ()
main = do
  _ <- getLine
  s <- getLine
  putStrLn $ solve s
