"""

"""

def main():

    path = sys.argv[1] if len(sys.argv) > 1 else "test.ppm"

    im = cv2.imread(path)
    cv2.imwrite(f"{path.split('.')[0]}.png", im)

    return


if __name__=="__main__":
    import os 
    import sys 
    # Change working directory to file directory for debbugging
    abspath = os.path.abspath(__file__)
    dname = os.path.dirname(abspath)
    os.chdir(dname)
    import cv2 
    main()