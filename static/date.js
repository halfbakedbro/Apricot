function datetime(id){

    date = new Date();
    year = date.getFullYear();
    month = date.getMonth();
    months = new Array('Janurary', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December');
    d = date.getDate();
    day = date.getDay();
    days = new Array('Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday');
    h = date.getHours();

    if (h < 10){
        h = "0"+h;
    }
    m = date.getMinutes();
    if (m<10){
        m = "0"+m;
    }

    s = date.getSeconds();
    if (s < 10){
        s = "0"+s;
    }

    result = ''+days[day]+' '+months[month]+' '+d+' '+year+' '+h+':'+m+':'+s;
    document.getElementById(id).innerHTML = result;
    setTimeout('datetime("'+id+'");', '1000');
    return true;
}