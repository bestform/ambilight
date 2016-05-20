Gaming Ambilight
================

This project is set up to watch a folder and from every jpeg that is written to that folder, create an average light color and send it
to a configured Philips Hue.

The plan is to watch a screenshot folder and let an addon populate the folder with a new screenshot every x seconds.

At this point, neither the addon, nor deleting the screenshots after they have been used, is implemented yet.
BUT: you can already watch the folder and take screenshots manually. Every screenshot will change the light of your lamp to the current average color in your game.

To do so, start the tool with the following flags:

--path=PATH_TO_YOUR_SCREENSHOT_FOLDER
--username=USERNAME_TO_YOUR_HUE_NETWORK
--ip=IP_OF_YOUR_HUE_NETWORK

Then just play the game and take a screenshot whenever you feel like it

If you want to delete every image after it has been processed, just add the option

--delete

Vendored dependencies
---------------------

All dependencies that are not part of the standard library are vendored using govendor (https://github.com/kardianos/govendor)
They are already part of the repository, but to use them, you either need to use at least go1.6 or - if you are on go1.5 - please set the environment variable GO15VENDOREXPERIMENT to 1.
