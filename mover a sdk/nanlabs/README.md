# The Space-x Challenge

The space-x team is designing their next launch to the international space station, they are recluting a group of the elite devs around the world and thought that you are gonna be a good fit.

Preparations are needed and they want to start organizing their tasks management so they’ve encoment you with your first task. The developer team uses Trello as their task manager app, but their management team (the one that creates the tasks) don’t want to use it, it’s too complicated for them. Is your job to create a bridge between these two teams.

The management team wants an endpoint that they can use to create the tasks, there are 3 flavors this day, but this could change in the future. A task may be:

1.  **An issue:** This represents a business feature that needs implementation, they will provide a short title and a description. All issues gets added to the “To Do” list as unassigned
2.  **A bug:** This represents a problem that needs fixing. They will only provide a description, the title needs to be randomized with the following pattern: bug-{word}-{number}. It doesn't matter that they repeat internally. The bugs should be assigned to a random member of the board and have the “Bug” label.
3.  **A task:** This represents some manual work that needs to be done. It will count with just a title and a category (Maintenance, Research, or Test) each corresponding to a label in trello. 

You need to create a post endpoint that will receive the tasks definition form the management team and create the corresponding cards in Trello [API Introduction](https://developer.atlassian.com/cloud/trello/guides/rest-api/api-introduction/) for the team to work with. Here are some examples:

**Expected deliverables:**

Code must be uploaded in a git repository (preferably GitHub).

We need instructions on how to execute the application.

Good luck, have fun!


```json
{
   "type":"issue",
   "title":"Send message",
   "description":"Let pilots send messages to Central"
}



Get Baord
curl --request GET --url \
'https://api.trello.com/1/boards/63c4903f1b996f00571f777b?key=64f8e58c56392537750ddb333e2ed257&token=ATTA0f11e8604137ea1b2222b196c12e50a175ced7a86cabdd0a0c9722f9b22bb916DF28BB37' \
--header 'Accept: application/json'


Get List
curl --request GET \
--url 'https://api.trello.com/1/lists/63c490512d8e4301239ceaa0?key=64f8e58c56392537750ddb333e2ed257&token=ATTA0f11e8604137ea1b2222b196c12e50a175ced7a86cabdd0a0c9722f9b22bb916DF28BB37'


Get Card
curl --request GET \
--url 'https://api.trello.com/1/cards/63cff19f08bca37420f1857b?key=64f8e58c56392537750ddb333e2ed257&token=ATTA0f11e8604137ea1b2222b196c12e50a175ced7a86cabdd0a0c9722f9b22bb916DF28BB37' \
--header 'Accept: application/json'

- Create Card




https://trello.com/1/authorize?expiration=never&scope=read,write,account&response_type=token&name=ServerToken&key=123xyz


curl --request GET \
  --url 'https://api.trello.com/1/lists/G0tzEZDj?key=64f8e58c56392537750ddb333e2ed257&token=ATTA0f11e8604137ea1b2222b196c12e50a175ced7a86cabdd0a0c9722f9b22bb916DF28BB37'


  curl --request GET \
  --url 'https://trello.com/b/G0tzEZDj/nanlabs?key=64f8e58c56392537750ddb333e2ed257&token=ATTA0f11e8604137ea1b2222b196c12e50a175ced7a86cabdd0a0c9722f9b22bb916DF28BB37'

https://trello.com/1/authorize?expiration=never&scope=read,write,account&response_type=token&name=ServerToken&key=64f8e58c56392537750ddb333e2ed257


 https://api.trello.com/1/cards?key=[apiKey]&token=[apiToken]&idList=[listID]&name='.urlencode ( ['ToDo']).'&desc='.urlencode (['cardDesc']);





curl --request DELETE \
  --url 'https://api.trello.com/1/cards/{id}/idMembers/{idMember}?key=APIKey&token=APIToken'



curl --request GET \
--url 'https://api.trello.com/1/cards/63dd84457300501f44802213/members?key=64f8e58c56392537750ddb333e2ed257&token=ATTA0f11e8604137ea1b2222b196c12e50a175ced7a86cabdd0a0c9722f9b22bb916DF28BB37' \
--header 'Accept: application/json'



curl --request GET \
--url 'https://api.trello.com/1/boards/63c4903f1b996f00571f777b/members?key=64f8e58c56392537750ddb333e2ed257&token=ATTA0f11e8604137ea1b2222b196c12e50a175ced7a86cabdd0a0c9722f9b22bb916DF28BB37'


curl --request POST \
  --url 'https://api.trello.com/1/labels?name={lala}&color=blue&idBoard=63c4903f1b996f00571f777b&key=64f8e58c56392537750ddb333e2ed257&token=ATTA0f11e8604137ea1b2222b196c12e50a175ced7a86cabdd0a0c9722f9b22bb916DF28BB37'


curl --request GET \
  --url 'https://api.trello.com/1/labels/63e40f80b1c831ec59bd0d5c?key=64f8e58c56392537750ddb333e2ed257&token=ATTA0f11e8604137ea1b2222b196c12e50a175ced7a86cabdd0a0c9722f9b22bb916DF28BB37'


Probar key y token.
probar recibir datos de la api en la app.


{
	"UUID": "1",
	"type": "issue",
	"title": "one",
	"description": "two",each corresponding to a label in trel
	"category": "three"
}

testing
mocks
evironment vars
logging
errors
docker
docker-compose
makefile
map db
readme



{
"id": "63e40f80b1c831ec59bd0d5c",
"idBoard": "63c4903f1b996f00571f777b",
"name": "Maintenance",
"color": "sky"
},
{
"id": "63e40f91208e91c4f5cbb5ac",
"idBoard": "63c4903f1b996f00571f777b",
"name": "Research",
"color": "red"
},
{
"id": "63e40f9dcc75976d321b4e7e",
"idBoard": "63c4903f1b996f00571f777b",
"name": "Test",
"color": "green"
},





curl --request POST \
  --url 'https://api.trello.com/1/cards?idList=63c490512d8e4301239ceaa0&key=64f8e58c56392537750ddb333e2ed257&token=ATTA0f11e8604137ea1b2222b196c12e50a175ced7a86cabdd0a0c9722f9b22bb916DF28BB37&name=koko&idLabels[]=63e40f9dcc75976d321b4e7e' \
  --header 'Accept: application/json'