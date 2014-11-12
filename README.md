Dead simple lib that provide a way to test fs like objects (which implement
os.File like interface).

Usage
=====

Just create new instance using `mockfile.New("somename")`, and then use common
file methods on it.
