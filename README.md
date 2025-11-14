## Monopoly simulation
Simulates a 4-player game of Monopoly played for a specified number of turns and counts total visits for each square. Results are then saved to a CSV file.
The goal of the simulation is to determine whether some properties are visited more frequently than others and are therefore more profitable to purchase.

Assumes players can always roll again if both rolled dice land on the same value.
Additionally some non-purchasable squares have been omitted from the results.

## Result chart
<img width="594" height="370" alt="image" src="https://github.com/user-attachments/assets/5a51818d-7c2e-42d3-89c5-6b7bc55a076f" />

The x-axis represents Monopoly board squares (starting from Go! as square 0). The values on the y-axis represent how many total visits each square got during the simulation.

Some squares have been omitted from the chart: the Go square, both jail squares, community chest and chance squares (their effects are still accounted for if they end up changing player's position)
Chart columns have been color-coded for easier reading according to their colors in the author's edition of Monopoly 
