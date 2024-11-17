<?php

$nilMap = [];
echo count($nilMap) . PHP_EOL;
// echo $nilMap['abc']; // PHP Warning:  Undefined array key "abc"
$nilMap['abc'] = 'def'; // 書き込み可能

$teams = [
    'giants' => ['坂本', '岡本'],
    'tigers' => ['糸井', '大山'],
    'carp'   => ['鈴木', '安部'],
];
echo $teams['giants'][0] . PHP_EOL; // 坂本

$user = [
    'name' => 'Taro',
    'age' => 25,
];
echo $user['name'] . PHP_EOL; // Taro
unset($user['age']);
var_export($user);
echo PHP_EOL;

// 連想配列のキーの存在チェック
echo isset($user['name']) ? 'true' : 'false'; // true
echo PHP_EOL;
echo isset($user['age']) ? 'true' : 'false'; // false
echo PHP_EOL;
