convert :: String -> [Int]
convert s = map (read :: String -> Int) $ words s

helper :: [Int] -> [Int] -> [Int]
helper acc []     = acc
helper acc (_:[]) = acc
helper acc xs     = helper (acc ++ [length $ xs']) xs'
    where
      xs' = filter (>0) $ map (subtract $ minimum xs) xs


solve :: String -> [Int]
solve s = result
    where
      ls     = convert s
      result = filter (>0) $ (length ls) : helper [] ls

main :: IO ()
main = do
  _ <- getLine
  s <- getLine
  mapM_ print $ solve s
