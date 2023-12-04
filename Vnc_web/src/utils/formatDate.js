export function formatTime(date) {
    const format = 'YYYY年MM月DD日 hh:mm:ss'
    const padZero = (num) => (num < 10 ? `0${num}` : num);

    const time = new Date(date);
    const year = time.getFullYear();
    const month = padZero(time.getMonth() + 1);
    const day = padZero(time.getDate());
    const hours = padZero(time.getHours());
    const minutes = padZero(time.getMinutes());
    const seconds = padZero(time.getSeconds());

    return format
        .replace('YYYY', year)
        .replace('MM', month)
        .replace('DD', day)
        .replace('hh', hours)
        .replace('mm', minutes)
        .replace('ss', seconds);
}
