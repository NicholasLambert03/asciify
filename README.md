# asciify 





An image to ascii art command line interface, written in go

## Table of Contents
- [Implementation](#implementation)
- [How to Use](#how-to-use)
- [Credits](#credits)

## Implementation
This project was written in go, using the cobra cli package. Images are read from filepath, supported image types are .jpg .gif and .png.  
A grayscale map of the image is created and used as a density function to map pixel intensity to character shape, then the option for adding colour to these characters, using the pixels original colour is available with the --color (-c) flag
Resizing of the image is also possible to allow it to fit better in the terminal

## How to Use
To use, download the repository and in its main directory run
```console
nicholas@easteregg:~Documents/asciify$  go build && go install
```
https://github.com/user-attachments/assets/dd13fd74-2267-45a5-8d3d-cbe8ec189a2b

Then when it is setup run 'asciify  --help'

```console
nicholas@easteregg:~$ aciify --help
Asciify is a command-line tool that converts .png .jpeg or .gif into ASCII art.

Usage:
  asciify [filepath] [flags]

Flags:
  -c, --color       Enable color output for the ASCII art using ANSI colours
  -h, --help        help for asciify
  -w, --width int   Custom width of the output ASCII art in characters
nicholas@Cube:~$ 
```

Then you are ready to play!
```console
                    
nicholas@Cube:~$ asciify Pictures/testimages/cat.png --width 40
                                        
                                        
                                        
                                        
       JJJ                    JJ        
       JJJ                    JJ        
       JJJ                    JJ        
       JJJJJ               JJJJJ        
       JJaaJ               JJaaJ        
       JJaaJJJJ          JJJJaaJ        
       JJaaJJJJ          JJJJaaJ        
       JJaaaJJJ          JJaaaaJ        
       JJaaaJJJ{{JJ{{JJ{{JJaaaaJ        
       JJaaaJJJ{{JJ{{JJ{{JJaaaaJ        
       JJaaJJJJ{{JJ{{JJ{{JJJJaaJ        
       JJaaJJJJ{{JJ{{JJ{{JJJJaaJ        
       JJJJJJJJJJJJJJJJJJJJJJJJJ        
     JJJJJJJJJJJJJJJJJJJJJJJJJJJJJJ     
     JJJJJJJMMMJJJJJJJJJMMMJJJJJJJJ     
     JJJJJJJMMMJJJJJJJJJMMMJJJJJJJJ     
     JJJJJMM  MMMJJJJJMMM  MMJJJJJJ     
     JJJJJMM  MMMJJJJJMMM  MMJJJJJJ     
     JJJJJJJMMMJJJJJJJJJMMMJJJJJJJJ     
     JJJJJJJJJJJJJJJJJJJJJJJJJJJJJJ     
     JJJJJJJJJJJJJJJJJJJJJJJJJJJJJJ     
     JJJJJJJJJJJJJJJJJJJJJJJJJJJJJJ     
  {{{{{{{{{{JJJJJ{{{{{JJJJJJ{{{{{{{{{{  
       JJJJJJJJJJJJ{{JJJJJJJJJJJ        
       JJJJJJJJJJJJ{{JJJJJJJJJJJ        
       JJJ{{JJJJJJJ{{JJJJJJJ{{JJ        
          {{JJJJJJJ{{JJJJJJJ{{          
        {{JJJJJJJ{{JJ{JJJJJJJJ{{        
       {    {{{{{JJJJJ{{{{{     {       
       {    {{{{{JJJJJ{{{{{     {       
     {{        JJJJJJJJJJ        {{     
     {{        JJJJJJJJJJ        {{     
                                        
                                        
                            
```

# Credits
Thank you to the team behind the ![Cobra CLI](https://github.com/spf13/cobra) for making an easy to implement command line interface framework






