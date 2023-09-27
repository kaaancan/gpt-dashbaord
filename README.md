# GPT Dashboard

My main motivation behind this project is that i want a very simple way of logging requests/responses to OpenAI and it should be self hostable for privacy reasons.

The v0 of this Project is WIP and should consist of the following basic features:

1. Proxy server which saves all requests and responses into a DB
2. Beautiful Dashboard so you can monitor your calls 

The main focus will be around the chat completion API so you can see your prompts and also your responses.

Also you should be able to link your OpenAI Call to any business logic you have so the proxy server should also add a callId which you can persist in your DB.