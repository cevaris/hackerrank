import Data.List
import Data.Char

lower  = map toLower
strip  = filter (/=' ')
helper = (== 26) . length . group . sort . lower . strip

solve :: String -> String
solve s = if helper s
          then "pangram"
          else "not pangram"

main :: IO ()
main = do
  input <- getLine
  (putStrLn . solve) input
