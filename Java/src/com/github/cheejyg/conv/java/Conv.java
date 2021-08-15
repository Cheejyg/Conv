package com.github.cheejyg.conv.java;

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import com.github.cheejyg.conv.java.*;
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * @author Cheejyg
 * @see Conv
 */
public class Conv {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Boolean conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * BooleanToInt parses <code>boolean</code> boolean_ to <code>int</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>int</code>
	 * @see Conv
	 */
	public static int BooleanToInt(boolean boolean_) {
		if (boolean_ == false) {
			return 0;
		} else if (boolean_ == true) {
			return 1;
		} else {
		}

		return 0;
	}

	/**
	 * BooleanToBoolean returns <code>boolean</code> boolean_.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be returned.
	 * @return <code>boolean</code>
	 * @see Conv
	 */
	public static boolean BooleanToBoolean(boolean boolean_) {
		return boolean_;
	}
}
