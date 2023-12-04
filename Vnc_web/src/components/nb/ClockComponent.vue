<script setup>
import {ref} from "vue";

const  date = ref("")
const  time = ref("")
var week = ['星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六'];
var timerID = setInterval(updateTime, 1000);
updateTime();
function updateTime() {
  var cd = new Date();
  time.value = zeroPadding(cd.getHours(), 2) + ':' + zeroPadding(cd.getMinutes(), 2) + ':' + zeroPadding(cd.getSeconds(), 2);
  date.value = zeroPadding(cd.getFullYear(), 4) + '-' + zeroPadding(cd.getMonth()+1, 2) + '-' + zeroPadding(cd.getDate(), 2) + ' ' + week[cd.getDay()];
};

function zeroPadding(num, digit) {
  var zero = '';
  for(var i = 0; i < digit; i++) {
    zero += '0';
  }
  return (zero + num).slice(-digit);
}
</script>

<template>
  <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
  <html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>时钟</title>

  </head>
  <body>

  <div class="clock">
    <p class="date">{{ date }}</p>
    <p class="time">{{ time }}</p>
  </div>



  </body>
  </html>
</template>

<style>
@charset "utf-8";
@import url("https://fonts.googleapis.com/css?family=Share+Tech+Mono");
html, body {
  height: 100%;
}

/*body {
  background: #0f3854;
  background: radial-gradient(ellipse at center, #0a2e38 0%, #000000 70%);
  background-size: 100%;
}*/

p {
  margin: 0;
  padding: 0;
}

clock {
  font-family: "Share Tech Mono", monospace;
  color: #ffffff;
  text-align: center;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  color: #daf6ff;
  text-shadow: 0 0 20px #0aafe6, 0 0 20px rgba(10, 175, 230, 0);
}
 .time {
  letter-spacing: 0.05em;
  font-size: 80px;
  padding: 5px 0;
}
.date {
  letter-spacing: 0.1em;
  font-size: 24px;
}
</style>