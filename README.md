# go-pub-sub

A simple PubSub library written in Golang.

Check out the [Developer Diary](./DEV_DIARY.md) to learn more of how I've developed this little project.

## Roadmap

- [x] Create initial architecture sketch.
- [x] Create initial package without _generics_ or use of files.
- [x] Add file management.
- [ ] Add Generics.

## Usage

First you must create a Pub/Sub agent:

```go
pubSub := pubsub.CreatePubSub()
```

After that you can create queues, subscribe and publish string data:

```go
pubSub.CreateQueue("my_queue")

sub, err := pubSub.Subscribe("my_queue")

err = pubSub.Publish("my_queue", message)
```

You can also add a logger to the agent, which will create a logging file for each queue with timestamps:

```go
err := pubSub.AddLogger()
```

Check out the [examples](cmd/) to learn more.