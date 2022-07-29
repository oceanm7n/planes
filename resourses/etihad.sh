#/bin/sh


bearer="Bearer T1RLAQJDmFRVL84iCqkJBI5YF6O1XPDcd0NxbJ/skxDwKKaU9hCH33wU1nAK6JbX408C0LbpAADQJJ8M058T7j2Jbesr2AKyTTwzfTfKhHdUTozZtyC+kfdUsRd/ulPMW2zpjUgFSceXWSgOynZ+vX9z1+9RBH7F8BFM9u1uAqTVzd4Bo8QQ16Y9c9wY3tEeBT1wLKy+uCVhushi1aEMpsFpYMYX9AgOHKh416IxOkN2EYDIVWfsDIah1/L2YkukslkzFpaOO1HnADUi1yQVv1h3p/e+ow92rIEK2FqiAfT5TXJjdpuy1BFVWayYQYdBaPT3OmXksb3pQZ1LNmodfhqq9bnLlwZqkA**"
EXECGUID="05370920-1d11-4e74-9445-8d2fcf6de546"



JIPCC="EYDX"

SRC="SVO"
DST="AUH"


MAXPRICENOTIFY=800
REQSLEEP=10
MAXWIDTH=80

DOTELEGRAM=true

#DDATE=2022-06-01



col=0

pecho () {

	[ $((col++)) -ge $MAXWIDTH ] && echo "" && col=1

	echo -n "$*"
}



[ "$1" == "" ] && echo usage $0 'YYYY-MM-DD <TO>' && exit 1
DDATE=$1
shift

[ "$1" != "" ] && DST=$1


mkdir -p /tmp/etihad
tmpfile=/tmp/etihad/$DDATE.$RANDOM.json

alltmpfiles="/tmp/etihad/$DDATE.*.json"
rm -rf $alltmpfiles


aedrate=`curl -s  'https://api.apilayer.com/exchangerates_data/convert?to=USD&from=AED&amount=1' --header 'apikey: 3yv7MhOBJHNi1UayxQrmc6iLzncX9SwF' | jq .result 2> /dev/null`
[ "$aedrate" == "" ] && echo Cannot get AED rate && exit 2
echo "1 AED = $aedrate USD"

echo Looking tickets from $SRC to $DST for $DDATE:

while true
do

	rm -rf $tmpfile

	curl -s -o $tmpfile \
	"https://dc.etihad.com/v4.3/dc/products/air/search?execution=$EXECGUID&jipcc=$JIPCC" \
	-H 'authority: dc.etihad.com' \
	-H 'accept: application/json' \
	-H 'accept-language: en-US,en;q=0.9,ru-RU;q=0.8,ru;q=0.7,he;q=0.6,de;q=0.5' \
	-H "authorization: $bearer" \
	-H 'content-type: application/json' \
	-H 'conversation-id: cl4eb78dwjow0szsqk0k1h6q5' \
	-H 'origin: https://dxbooking.etihad.com' \
	-H 'referer: https://dxbooking.etihad.com/' \
	-H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36' \
	--data '{"cabinClass":"Economy","awardBooking":"false","promoCodes":[],"searchType":"BRANDED","itineraryParts":[{"from":{"useNearbyLocations":false,"code":"'$SRC'"},"to":{"useNearbyLocations":false,"code":"'$DST'"},"when":{"date":"'"$DDATE"'"}}],"passengers":{"ADT":1}}'   --compressed

	if [ $? -eq 0 ]
	then
		curpr=`cat $tmpfile | jq -rc '.unbundledOffers[] | .[] | .seatsRemaining.count,.total.alternatives[][0].amount ' 2> /dev/null | head -n 2 | paste -sd"\t"`

		[ $? -ne 0 ] && cp -af $tmpfile $tmpfile.error1

		if [ "$curpr" != "" ]
		then

			oldpricenum=$pricenum
			[ "$oldpricenum" == "" ] && oldpricenum=0


			numtic=`echo $curpr | cut -f1 -d' '`

			price=`echo $curpr | cut -f2 -d' '`
			priceusd=`echo $price*$aedrate | bc -l`
			pricenum="${priceusd%.*}"

			curpr="$numtic"" ""$pricenum"


			if [ "$oldpr" == "" ]
			then

				oldpr="$curpr"
				echo -e "\n"`date "+%Y-%m-%d %T"`"\t""$curpr"
				col=0

				#[ "$DOTELEGRAM" == "true" ] && [ "$pricenum" != "" ] && [ $pricenum -le $MAXPRICENOTIFY ] && /home/arami/telegram.sh  --loud "$DDATE"": ""$curpr"
				[ "$DOTELEGRAM" == "true" ] && [ "$pricenum" != "" ] && /home/arami/telegram.sh  --loud "Start monitoring $DDATE"": ""$curpr"
			else
				if [ "$oldpr" != "$curpr" ]
				then
					oldpr="$curpr"
					echo -e "\n"`date "+%Y-%m-%d %T"`"\t""$curpr"
					col=0

					[ "$DOTELEGRAM" == "true" ] && [ "$pricenum" != "" ] && ( [ $pricenum -le $MAXPRICENOTIFY ] || ( [ $oldpricenum -le $MAXPRICENOTIFY ] && [ $pricenum -gt $oldpricenum ] ) ) && /home/arami/telegram.sh  --loud "$DDATE"": ""$curpr"

				else
					pecho "."
				fi
			fi
		else
			mess=`cat $tmpfile | jq -rc '.messages[].code' 2> /dev/null` 

			if [ "$mess" == "flight_not_found.use_alternate_search" ]
			then
				pecho "n"
			else
				mess=`cat $tmpfile | jq -rc '.message' 2> /dev/null`
				if [ "$mess" == "Authentication failed due to invalid credentials" ]
				then
					echo "C"
					echo "$mess".
					break
				else
					[ $? -ne 0 ] && cp -af $tmpfile $tmpfile.error2
					pecho "e"
				fi
			fi
			

		fi
	else
		pecho "E"
	fi


	sleep $REQSLEEP

done




