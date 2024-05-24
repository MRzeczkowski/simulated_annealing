# Simulated Annealing

This is a simple project that uses simulated annealing to find the minimal value of the Rastrigin function.
Simulated annealing is a probabilistic technique for approximating the global optimum of a given function. It is particularly useful for large optimization problems.
I've used the Cauchy distribution for mutations with Gamma = 0.05 since I had good results with it in my previous projects.

## Results
Below is the output of the program. The table shows the best parameters found for each cooling scheme after running multiple tests. The average result represents the average Rastrigin function value achieved using the best parameters.

Simulation Parameters:
- Number of dimensions: 3
- Max iterations per test: 1000
- Number of tests per cooling scheme: 10

Best parameters for each cooling scheme
| Cooling scheme | Initial temperature | Min temperature | Cooling rate | Max iterations at temperature level | Average result |
|-|-|-|-|-|-|
| logarithmic | 500.00 | 0.80 | None | 175 | 10.7017 |
| harmonic | 1400.00 | 0.10 | None | 25 | 0.6977 |
| geometric | 900.00 | 0.10 | 0.80 | 25 | 2.8971 |
| linear | 800.00 | 0.10 | 0.95 | 200 | 12.7457 |
| exponential | 800.00 | 0.40 | None | 100 | 1.7653 |

### Cooling Schemes:
- **Geometric**: The temperature is multiplied by a constant factor (cooling rate) at each step.
- **Linear**: The temperature is decreased by a constant amount at each step.
- **Exponential**: The temperature decreases according to an exponential schedule.
- **Logarithmic**: The temperature decreases according to a logarithmic schedule.
- **Harmonic**: The temperature follows a harmonic series decrease.

Note that logarithmic, harmonic and exponential cooling schemes don't have a configurable cooling rate.

### Conclusion:
The best cooling scheme found in my experiments is the harmonic scheme, which produced the lowest average result for the Rastrigin function. This indicates that the harmonic cooling schedule is particularly effective for this type of optimization problem.

Other notable observations:
- The harmonic scheme had the highest initial temperature but required fewer iterations at each temperature level, indicating faster convergence.
- The geometric and exponential schemes also performed well, achieving relatively low average results.
- The logarithmic and linear schemes had higher average results, suggesting they may be less effective for this specific optimization problem.
