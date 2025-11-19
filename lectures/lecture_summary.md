
## linear type
c := make(chan1.1 int) 
## usage type
c := make(chan1+1 int)
## session type
?

### typing environment split to read and write: 
linear type: {c -> chan1.1 int} = {c -> chan1.0 int} + {c -> chan0.1 int}
usage type: {c -> chan0+0 int} = {c -> chan1.0 int} + {c -> chan0.1 int}

