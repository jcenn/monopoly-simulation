## Monopoly simulation
Simulates a 4-player game of Monopoly played for a set number of turns and counts total visits for each square. Results are then written to a CSV file.
The aim of this simulation was to find out if some squares are visited more often than others and thus are more profitable if purchased.

Assumes players can always roll again if both rolled dice land on the same value.
Additionally some non-purchasable squares have been omitted from the results.

## Result chart
<img width="594" height="370" alt="image" src="https://github.com/user-attachments/assets/5a51818d-7c2e-42d3-89c5-6b7bc55a076f" />

The x-axis represents Monopoly squares (starting from Go! being 0). Values on the y-axis represent how many visits a certain square got during the simulation.

Some squares have been ommited from the chart, namely: the Go square, both jail squares and community chest and chance squares (their effects are still accounted for if they end up changing player's position)
Chart columns for properties have been color-coded for easier reading according to their colors in author's edition of Monopoly 
