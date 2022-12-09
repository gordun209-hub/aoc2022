module Main where
import           Data.List
import           Data.Map.Strict (Map)
import           Data.String
import           System.IO

-- day 1 aoc
-- move 1 from 8 to 1


-- day 1 advent of code 2023

-- input input = '1000 \
-- 2000 \
-- 3000 \
--
-- 4000
--
-- 5000
-- 6000
--
-- 7000
-- 8000
-- 9000
--

-- 10000

square x = x * x

sum_of_squares x y = square x + square y

abslt :: (Ord a, Num a) => a -> a
abslt x
     |x < 0 = -x
     |x > 0 = x
     |x == 0 = 0

sqrt_iter :: (Ord a, Fractional a) => a -> a -> a
sqrt_iter guess x
          | good_enough guess x = guess
          | otherwise = sqrt_iter (improve guess x) x where
                improve guess x = (guess + x / guess) / 2
                good_enough guess x = abslt (square guess - x) < 0.001


f a = sum_of_squares (a + 1) (a * 2)
main = do
    print  (sqrt_iter 5 2)
    print ((2^) $ (*30) 0)
