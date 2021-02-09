Tool for collecting git log data.  
Inspired by https://didoo.medium.com/measuring-the-impact-of-a-design-system-7f925af090f7

Script takes one argument, git repository url

Output a json file with stats from each commit  
Values:
 - CommitDate
 - NumFileChanged
 - NumInsertions
 - NumDeletions

 Output logs is a bit buggy!