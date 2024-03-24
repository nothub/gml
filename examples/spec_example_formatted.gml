<gml spec="1.0">

    <tag>

        <header>

            <!-- how, who, what and where -->
            <client>

                <!-- application name -->
                <name>Laser Tag</name>

                <!-- application version -->
                <version>2.0</version>

                <!-- user name on 000000book.com, optional -->
                <username>MyUserName</username>

                <!-- URL to .gml data on 000000book.com, optional -->
                <permalink>http://000000book.com/data/156/</permalink>

                <!-- comma-separated -->
                <keywords>katsu,paris,2010</keywords>

                <!-- iPhone uuid, MAC address, etc -->
                <uniquekey>28sks922ks992</uniquekey>

                <ip>192.168.1.1</ip>

                <!-- unixtime -->
                <time>1928372722</time>

                <location>
                    <lon>-39.392922</lon>
                    <lat>53.29292</lat>
                </location>

            </client>

            <!-- This is all stuff that relates to the orientation and dimensions of the client -->
            <!-- So that we know how to re-map the 0.0-1.0 coordinates that come in for each point -->
            <!-- Also for figuring out the down vector for devices with accelerometers and how that effects drips -->
            <!-- All numbers should be between 0.0 - 1.0 -->
            <environment>
                <offset>
                    <x>0.0</x>
                    <y>0.0</y>
                    <z>0.0</z>
                </offset>
                <rotation>
                    <x>0.0</x>
                    <y>0.0</y>
                    <z>0.0</z>
                </rotation>
                <up>
                    <x>0.0</x>          <!-- commonly up for iphone apps -->
                    <y>-1.0</y>         <!-- most common -->
                    <z>0.0</z>
                </up>
                <screenbounds>     <!-- use this as your multipler to get 0.0 to 1.0 back to right size - pts should never go off 0.0 to 1.0 -->
                    <x>1024</x>
                    <y>768</y>
                    <z>0</z>
                </screenbounds>
                <origin>
                    <x>0</x>
                    <y>0</y>
                    <z>0</z>
                </origin>
                <realscale>     <!-- how these units relate to real world units - good for laser tag -->
                    <x>1000</x>
                    <y>600</y>
                    <z>0</z>
                    <unit>cm</unit>
                </realscale>
                <audio>youraudio.mp3</audio>                        <!-- path to audio file -->
                <background>yourimage.jpg</background>     <!-- path to image file -->
            </environment>

        </header>

        <drawing>
            <!-- for all stroke and movement stuff it helps to have everything inside the stroke tag -->
            <!-- this way it is easy to get a sense of order to events -->

            <stroke isdrawing="false">     <!-- for non drawing mouse movements -->
                <pt>
                    <x>0.0</x>
                    <y>0.0</y>
                    <z>0.0</z>         <!--this is optional -->
                    <t>0.013</t>       <!-- time is optional too -->
                    <!-- NOTE: older versions of GML use <time> instead of <t> -->
                </pt>
            </stroke>

            <stroke>                                 <!-- by default stroke drawing is true -->

                <!-- each stroke could be drawn with a different brush -->
                <!-- if no brush tag is found for a stroke then it inherits the previous settings -->
                <brush>
                    <mode>0</mode>     <!-- same as uniqueStyleID but an internal reference -->
                    <uniquestyleid>LaserTagArrowLetters</uniquestyleid> <!-- unique blackbook string for your style -->
                    <!-- see note about spec at the bottom - like unique style but with extra info -->
                    <spec>http://aurltodescribethebrushspec.com/someSpec.xml</spec>
                    <width>10</width>
                    <speedtowidthratio>1.5</speedtowidthratio>     <!-- put 0 for fixed width -->
                    <dripamnt>1.0</dripamnt>
                    <dripspeed>1.0</dripspeed>
                    <layerabsolute>0</layerabsolute>     <!--Think photoshop layers-->
                    <color>
                        <r>255</r>
                        <g>255</g>
                        <b>255</b>
                        <a>255</a>     <!-- optional -->
                    </color>
                    <dripvecrelativetoup>     <!-- what angle do our drips go in relation to our up vector -->
                        <x>0</x>
                        <y>1</y>
                        <z>0</z>
                    </dripvecrelativetoup>
                </brush>

                <pt>
                    <x>0.0</x>
                    <y>0.0</y>
                    <z>0.0</z>                     <!--this is optional -->
                    <t>0.013</t>       <!-- time is optional too -->
                </pt>

                <pt>
                    <x>0.0</x>
                    <y>0.0</y>
                    <z>0.0</z>                     <!--this is optional -->
                    <t>0.023</t>       <!-- time is optional too -->
                </pt>


            </stroke>

            <!-- this stroke inherits the previous stroke properties -->
            <!-- but changes color and draws on the layer below -->
            <stroke>
                <info>     <!-- optional info - more stuff soon-->
                    <curved>true</curved>
                </info>
                <brush>
                    <color>
                        <r>255</r>
                        <g>255</g>
                        <b>0</b>
                    </color>
                    <layerrelative>     <!-- this means one layer bellow the previous layer  -->
                        -1
                    </layerrelative>
                </brush>

                <pt>
                    <x>0.0</x>
                    <y>0.0</y>
                </pt>

                <pt>
                    <x>0.0</x>
                    <y>0.0</y>
                </pt>

            </stroke>

            <stroke>
                <pt>
                    <pres>0.5</pres>     <!-- Optional. Preasure range from 0 to 1 -->
                    <rot>0.5</rot>     <!-- Optional. Rotation range from 0 to 1 for 0 to 2*PI -->
                    <dir>                  <!-- Optional Direction -->
                        <x></x>                  <!-- range from 0 to 1 -->
                        <y></y>                  <!-- range from 0 to 1 -->
                        <z></z>                  <!-- Optional inside direction. Range from 0 to 1 -->
                    </dir>
                </pt>
            </stroke>
        </drawing>

    </tag>

</gml>
