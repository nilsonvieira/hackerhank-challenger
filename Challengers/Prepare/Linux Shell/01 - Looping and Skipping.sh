for i in {1..99}
do
        returnMod=$(($i % 2))
        if [ "$returnMod" -ne "0" ]; then
                echo $i
        fi
done