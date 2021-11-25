# om-stream

This is intended as a rework of open-match. 
The original objective was to move from having a synchroniser singleton, 
to using a streaming database which should increase scope for capacity and resilience.

Other beneficial design changes:
1. In open-match Tickets can be submitted into the system that do not match any 
   existing match profiles. In om-stream a ticket is sent onto a stream specific 
   to the profile the ticket matches, if a ticket doesn't match a profile an error 
   is returned to the client.
2. In open-match frontend.WatchAssignments polls the database 
   (every 30 milliseconds by default) until the Assignment value is no longer empty.
   om-stream uses a publish-subscribe system to announce when tickets have been updated.
3. In open-match the director is a singleton that triggers when matchmaking occurs, 
   disseminates match profiles to the open-match backend, allocate a game server and  
   assign the game server connection string to the tickets. 

Below is a sequence diagram that gives an overview of how om-stream works
![uml](./docs/sequence.png)
