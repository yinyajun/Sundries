# How to write main function in Python?
import sys
import getopt 

class Usage(Exception):
  def __init__(self,msg):
    self.msg = msg
    
def main(argv = None):
  if argv is None:
    argv = sys.argv
  try:
    try:
      opts, args = getopt.getopt(argv[1:], "h", ["help"])
    except getopt.error, msg:
      raise Usage(msg)
      # more code
  except Usage, err:
    print>>sys.stderr, err.msg
    print>>sys.stderr, "for help use --help"
    return 2
  
if __name__="__main__":
  sys.exit(main())
