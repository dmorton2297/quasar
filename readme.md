# Goal - create an extendable relational data store service that

MVP
1. Has relational database for storing data for a strongly typed entity (relational)
2. Generates messages for the following events.
    1. Event was deleted
        1. Metadata - ID and blob that was deleted
    2. Event was created - ID and blob that was created
    3. Event was updated - ID and updated blob vs original
    4. event was read - ID and data that was read
3. Data passed in messages in compressed to save storage. 
    1. What we compare about is good compression ratio
    2. Decompression needs to be linear to reduce size


