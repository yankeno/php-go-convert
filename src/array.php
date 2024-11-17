<?php

$fruits = ['apple', 'banana', 'cherry'];
foreach ($fruits as $fruit) {
    echo $fruit . PHP_EOL; // apple, banana, cherry
}

$a = ['a', 'b'];
$b[0] = 'a';
$b[1] = 'b';
echo ($a === $b) ? 'true' : 'false'; // true
echo PHP_EOL;

$c[] = 'a'; // 要素の追加
var_export($c); // array ( 0 => 'a')
echo PHP_EOL;

$d = &$a; // 参照渡し
$d[0] = 'A';
var_export($a); // array ( 0 => 'A', 1 => 'b')
echo PHP_EOL;
