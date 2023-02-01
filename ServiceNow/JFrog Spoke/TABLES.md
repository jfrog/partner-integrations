# ServiceNow Tables for Xray

## Overview
The following tables are all included as a part of the JFrog Spoke.
Use this reference to navigate the tables within your flows.

## Tables

### Violations

The following is the violations table. It contains simple things like the watch and policy name that produced this violation.

<img src="./images/Violations.png?raw=true" width="800">

### Issues

The following is the issues table. It contains a reference to its corresponding violation. (A reference in ServiceNow is a reference to an associated record in another table.)

<img src="./images/Issues.png?raw=true" width="800">

### Impacted Artifacts

The following is the impacted artifacts table. It contains a reference to its corresponding issue.

<img src="./images/ImpactedArtifacts.png?raw=true" width="800">

Column Explanations:
- Path: Refers to the Repo Key of an impacted artifact.
- Display Name: Refers to the name of the impacted artifact.

### Infected Files

The following is the infected files table. It contains a reference to its corresponding impacted artifact.

<img src="./images/InfectedFiles.png?raw=true" width="800">

Column Explanations:
- Path: Refers to the repo key, or the path up until the display name of the infected file.
- Display Name: Refers to the name of the infected file.
