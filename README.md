# om-kafka

This is intended as a rework of open-match. 
The original objective was to move from having a synchroniser singleton, 
to using apache kafka as the stream manager which should increase scope 
for capacity and improve reliability.

Other beneficial design changes became apparent as this work proceeded:
1. In open-match Tickets can be submitted into the system that do not match any 
   existing match profiles. In om-kafka a ticket is sent onto a stream specific 
   to the profile the ticket matches, if a ticket doesn't match a profile an error 
   is returned to the client.
1. In open-match frontend.WatchAssignments polls the database 
   (every 30 milliseconds by default) until the Assignment value is no longer empty.
   We use a publish-subscribe system to announce when tickets have been updated.
