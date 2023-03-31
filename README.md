# vault-spike


# milestones 

### slow and controlled 
milestone 1:
the previous commit has wave 10 and wave 20 slowly rolling out each application with a sleep 1 second up front and a sleep 10 seconds at the tail end of 
each -components wrapper. no `PostSync` hooks 


milestone 2:
we removed the job that ran first and slept 10 seconds to start a sync wave because it proved to be useless.
we also removed `PostSync` hooks completely and relied only on sync waves 10,20,100
the final job in a sync wave seems to be what is holding up subsequent sync waves

milestone 3:
using raw manifests from the kubernetes-toolkit upstream to halt applications from moving forward.
we can successfully hold sync waves with these final jobs
