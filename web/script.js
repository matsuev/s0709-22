// Use WebSocket transport endpoint.
const client = new Centrifuge('ws://localhost:8080/centrifugo/connection/websocket');

const userChannel = client.newSubscription("vasya")
userChannel.on('publication', function (ctx)
{
   console.log(ctx);
});
userChannel.subscribe()


// Allocate Subscription to a channel.
const newsSub = client.newSubscription('news');

// React on `news` channel real-time publications.
newsSub.on('publication', function (ctx)
{
   console.log(ctx);
});

// Trigger subscribe process.
newsSub.subscribe();

// Trigger actual connection establishement.
client.connect();