START
   |
   v
Is len(os.Args) == 2 ?
   |             \
  No             Yes
   |               |
Print "a"         input = os.Args[1]
and EXIT            |
                    v
        Loop over each char in input
                    |
          Is char == 'a' ?
            |         \
           Yes        No (continue loop)
            |            |
      Print "a"          |
      and EXIT           |
                    (loop ends)
                    |
             Print blank line
                    |
                   END
