# quiz-1
使用技術：Golang、MongoDB、Docker<br>
1.先下載檔案下來後，在命令列輸入<b>「docker-compose up」</b><br>
啟動API服務容器以及MongoDB容器 <br>
![image](comment/img/1.png)

2.看到服務啟動成功，此時我們目前API啟動服務位址是
http://localhost:8080 <br>
![image](comment/img/2.png)

3.使用MongoDB資料庫 <br>
可以先在網路下載MongoDB Compass GUI https://www.mongodb.com/try/download/compass

4.MongoDB的連線位址 <b>「mongodb://localhost:27017/」</b><br>
![image](comment/img/3.png)

5.此時可以使用API服務，去建立comment 資料 <br>
此專案中的api：<br>

1.POST /comment - 創建評論 <br>
2.Get /comment/{uuid} - 取得單一評論 <br>
3.Put /comment/{uuid} - 更新單一評論 <br>
4.Delete /comment/{uuid} - 刪除單一評論 <br>

資料格式為 <br>
{ <br>
    "parentid": "", - string <br>
    "comment": "根據中央氣象局地震測報中心地震報告，這起規模...", - string <br>
    "author": "氣象局網站", - string  <br>
    "update": "2022-06-01T01:12:33Z", - string <br>
    "favorite": false - string - boolean <br>
} <br>

6.成功新增一筆資料後，可以前往剛剛的MongoDB Compass GUI 刷新，查看 mydb資料庫中的comment<br>
裡面就會顯示剛剛新增的資料<br>

![image](comment/img/4.png)
