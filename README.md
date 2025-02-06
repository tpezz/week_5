# Assignment 5 – AI Code Generation (MSDS 431)

This GitHub contains my work for Assignment 5 in MSDS 431. The goal of the assignment was to explore the use of free AI tools to assist in coding and to generate code automatically. I revisited my earlier “Go for Statistics” project (week 4) and evaluated the ability to read a CSV file and perform linear regression across three programming languages: Go, Python, and R.

---

## Project Overview

- **Objective:** Test both AI-assisted programming and AI-generated code in order to learn about how they can help developers. I also wanted to see how the code with these techniques compares to my origional code. 

- **Task:** Read a CSV file (containing the Anscombe Quartet data) and perform linear regression on the four x/y pairs.

---

## Plan and results

### Step 1. Revisiting "Go for Statistics"
I started with my original code from the "Go for Statistics" assignment. This code reads a CSV file and performs linear regression on the Anscombe Quartet dataset.

### Step 2. AI-Assisted Programming with GitHub Copilot
For each of the prgramming languages I used the free version of GitHub Copilot in VS Code extension on the origional file from that lanaguage with the following prompts:
**Prompt 1:**  
  "Review this file, make updates so that it is streamlined, readable, and aligns with idiomatic practices"
  
**Prompt 2:**  
  "Fix any bugs so the file runs well"
  
**Prompt 3:**  
  "Add comments to the code to help someone new understand it"

**Experience with Copilot:**
**Pros:**  
  - Helped simplify and streamline the logic in the code
  - Improved readability and added error handling and logging
  - Provided useful suggestions for all three languages
**Cons:**  
  - Sometimes it was really slow
  - When the same task was repeated in different languages, Copilot occasionally got confused and suggested unrelated files
  - It performed best with Python, while Go and R sometimes needed extra tweaks

### 3. AI-Generated Code with LLM Agents (ChatGPT 4o)
I also experimented with AI-generated code using ChatGPT 4o. My prompts included:
**Prompt 1:**  
  "You have an Anscombe quartet dataset in a file called `Anscombe_quartet_data.csv`. Create a file in {language} that reads the CSV and performs linear regression on the four x/y pairs. Keep the code as simple and straightforward as possible."
  
**Prompt 2:**  
  "Confirm that everything works, given that the CSV file is in the same directory"

**Results:**
- The AI-generated code worked well after some minor adjustments
- Python and R outputs were very streamlined
- The R code occasionally tried to use Python code for other languages, so I had to watchout for that

---

## Summary of Code Results by Language

### Go
- **Original Code:**  Functional but lacked some formatting and robust error handling
- **AI-Assisted Version:**  Improved error handling, streamlined logic, better logging, and clearer comments  
- **Fully AI-Generated Version:**  Simpler code but lost some timing capapbilites and robustness

### Python
- **Original Code:**  Provided a clear structure but could be more efficient
- **AI-Assisted Version:**   Retained the original intent while improving efficiency and readability
- **Fully AI-Generated Version:**  Riskier due to removal of key details. Could be used as a starting point

### R
- **Original Code:**  Worked but was not ideal (I dont code R very well)
- **AI-Assisted Version:**  Best option. It refined the code without compromising my intended functionality
- **Fully AI-Generated Version:**  Good for quick analysis but missed alot of key things I wanted

**Overall Best Approach:**  
**AI-Assisted Programming** provided the best balance between improving the original code and maintaining the features I wanted


## Resources

- **GitHub Copilot Documentation for VS Code:**  
  [GitHub Copilot Docs](https://docs.github.com/en/copilot)
  
- **ChatGPT 4o:**  
  [Overview of chatGPT 4o](https://openai.com/index/hello-gpt-4o/)
  
- **Anscombe Quartet Data Information:**  
  [Anscombe Quartet Data Kaggle Dataset](https://www.kaggle.com/datasets/carlmcbrideellis/data-anscombes-quartet)

---

## Included Folders
  - `FOLDER: AI Assisted Co Pilot` (All files where I used AI assistance)
  - `FOLDER: Original` (All original files from “Go for Statistics” assignment)  
  - `FOLDER: AI Generated` (All AI Generated files and chat logs)
  

## Conclusion and Recommendation

This assignment showed that AI tools are best used as assistants rather than complete replacements for human coding. The AI-assisted approach helped maintain the original intent and structure of the code while improving clarity, error handling, and efficiency. For startups, combining experienced programmers with AI tools can reduce workload and lead to higher quality software development.

I also did some testing with o3-mini-high which is ChatGPTs more advanced coding model but it is apart of their $20 a month tier. I would highly recommend it as it performed far better better than GPT 4o. 
