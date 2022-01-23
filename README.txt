  _______       _ _       _           _                 _        
 |__   __|     (_) |     | |         | |               | |       
    | |_      ___| |_ ___| |__ ______| |__   ___   ___ | | _____ 
    | \ \ /\ / / | __/ __| '_ \______| '_ \ / _ \ / _ \| |/ / __|
    | |\ V  V /| | || (__| | | |     | | | | (_) | (_) |   <\__ \
    |_| \_/\_/ |_|\__\___|_| |_|     |_| |_|\___/ \___/|_|\_\___/
by Unbewohnte

[Eng]
! Purpose
This program`s purpose is to send custom messages
to Vk`s users/group chats and discord`s text channels
via webhooks when the Twitch stream is online.

! First run
When you run it the first time - it`ll generate a config.cfg
in your working directory and exit.

! Config file
Config file`s contents are present in JSON format.

Field "TwitchName" must contain Twitch channel`s name. The
CLI will be checking if this user has started streaming or not.

In block "Keys" you should paste at least all Twitch keys and 
at least one of the fields that are left (Discord`s webhook url or Vk`s api key)

If "force-send" == true - the program won`t check for any streams and just 
send messages you`ve stated it to send.

In block "Messages" you can specify what kind of messages the CLI will 
send in case of the stream. Usually it`s a little opening followed by a
stream`s link.

!! "receiver_id" and "is_group_chat"
If you want to send a message to the group chat in VK and 
set "is_group_chat" to true, then "receiver_id" must be your
personal id of that group chat. If you navigate to it in your
browser - you`ll see such pattern in the end of the URL: c20. In this case
20 is my personal id for some of my group chats. You should
set yours to "receiver_id". 

! Next runs
After you`ve finished your configuration - the program`s ready to 
work. Just execute it again and it`ll check every 5 minutes for specified
user`s stream. If it`s started - it`ll send your custom messages and exit.


! Api keys 

!! Twitch 
https://dev.twitch.tv/console - here, create an app and grab all the necessary keys

!! Discord
Create a webhook in the server`s settings and copy its URL

!! Vk
Create your own app, or proceed here: https://vkhost.github.io/ and grab your API key without a headache 



[Ru]
! Назначение
Эта программа предназначена для отсылания оповещений о начавшемся
Twitch стриме на указанном канале пользователям/группе вКонтакте и 
в текстовый канал Discord через webhook   

! Первый запуск
При первом запуске cli генерирует файл в рабочей директории с именем
"config.cfg".

! Конфигурационный файл
Файл структурирован в виде JSON формата.

В поле "TwitchName" следует указать имя канала
на Твиче. Программа будет отслеживать начало стрима данного
пользователя.

В блоке "Keys" следует вставить как минимум все ключи от Twitch и
хотя-бы один из оставшихся url или ключей от ВК или Дискорда.

Если "force-send" == true - программа не станет ждать начала стрима
и сразу отошлёт указанные вами сообщения

В блоке "Messages" следует указать сообщение, которые вы хотите
отослать в случае начала стрима. Обычно это небольшое вступление 
и ссылка на стрим.

!! "receiver_id" и "is_group_chat"
Если вы хотите отослать сообщение в групповой чат ВК и 
выставили is_group_chat на true, то receiver_id будет ваше
личное ид данного чата. Чтобы его получить, зайдите в него и
в конце URL вы заметите что-то наподобие такого: с20. 20 в данном
случае и есть ид группового чата.  

! Последующие запуски
После настройки конфигурационного файла программа готова к 
полному использованию. Просто снова запустите её и она каждые
5 минут будет проверять наличие стрима на указанном канале. 
Если стрим идёт - она отправит сообщения и закроется.

! Api ключи 

!! Twitch 
https://dev.twitch.tv/console - создай своё приложение и забери все необходимые ключи

!! Discord
Создай вебхук в настройках сервера и скопируй ссылку

!! Vk
Создай своё приложение, или просто зайди сюда: https://vkhost.github.io/ и забери свой ключ без головной боли
