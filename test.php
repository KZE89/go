<?php
header('Content-Type: text/html; charset=utf-8');

$myCurl = curl_init();
curl_setopt_array($myCurl, array(
    CURLOPT_URL => 'http://localhost:8000/login',
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_POST => true,
    CURLOPT_POSTFIELDS => http_build_query(array('login' => 'test', 'pass' => '123456'))
));
$response = curl_exec($myCurl);


echo "Ответ на Ваш запрос: ".$response ."<br><br><br>";

echo curl_getinfo($myCurl, CURLINFO_HTTP_CODE)."<br>";
curl_close($myCurl);


$myCurl = curl_init();
curl_setopt_array($myCurl, array(
    CURLOPT_URL => 'http://localhost:8000/login/pass',
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_POST => true,
    CURLOPT_POSTFIELDS => http_build_query(array('login' => 'test', 'pass' => '123456', 'newPass' => '123456'))
));
$response = curl_exec($myCurl);


echo "Ответ на Ваш запрос: ".$response ."<br><br><br>";


echo curl_getinfo($myCurl, CURLINFO_HTTP_CODE)."<br>";
curl_close($myCurl);

$t = array('BigNumber' => 100, 'Text' => 'taza');

$myCurl = curl_init();
curl_setopt_array($myCurl, array(
    CURLOPT_URL => 'http://localhost:8000/login/dowork',
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_POST => true,
    CURLOPT_POSTFIELDS => http_build_query(array('login' => 'test', 'value' => json_encode($t)))
));
$response = curl_exec($myCurl);

echo json_encode($t)."<br><br><br>";
echo "Ответ на Ваш запрос: ".$response."<br><br><br>";


echo curl_getinfo($myCurl, CURLINFO_HTTP_CODE)."<br>";
curl_close($myCurl);
?>