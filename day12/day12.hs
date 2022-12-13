module Main where

main :: IO ()
main = do
    input <- readFile "../sample12"
    print (split input)

type Grid
     = [[Char]]

split :: String -> [String]
split [] = []
split (c:cs) | c == '\n'  = "" : rest
             | otherwise = (c : head rest) : tail rest
    where rest = split cs
