Spectator
===================

This is a bot that listens to IRC channels and logs what usernames are associated with what channels.
It also keeps track of who directly references whom.
It then stores that information in a locally running Neo4j server.

This allows community identification between individuals as well as overlap in channel membership.  Currently no NLP or semantic analysis is done, but it is an obvious next step in determining real communities and hubs/experts within them.
